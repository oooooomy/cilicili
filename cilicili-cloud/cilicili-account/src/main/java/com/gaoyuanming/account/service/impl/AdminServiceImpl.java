package com.gaoyuanming.account.service.impl;

import com.gaoyuanming.account.feign.EmailFeign;
import com.gaoyuanming.account.model.entity.Admin;
import com.gaoyuanming.account.repository.AdminRepository;
import com.gaoyuanming.account.service.AdminService;
import com.gaoyuanming.common.model.support.ResponseResult;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

@Service
public class AdminServiceImpl implements AdminService {

    @Resource
    private AdminRepository adminRepository;

    @Resource
    private EmailFeign emailFeign;

    @Override
    public Admin save(Admin admin) {
        return adminRepository.save(admin);
    }

    @Override
    public List<Admin> findAll() {
        return adminRepository.findAll();
    }

    @Override
    public void delete(String id) {
        adminRepository.deleteById(id);
    }

    @Override
    public Admin loginByPassword(Admin admin) throws Exception {
        Admin one = adminRepository.findAllByEmailAndPassword(admin.getEmail(), admin.getPassword());
        if (one == null) throw new Exception("密码错误");
        return one;
    }

    @Override
    public Admin loginByCode(String email, String code) throws Exception {
        Admin admin = adminRepository.findByEmail(email);
        if (admin == null) throw new Exception("身份验证失败");
        ResponseResult<Object> result = emailFeign.checkEmailCode(email, code);
        if (!result.isStatus()) throw new Exception(result.getMsg());
        return admin;
    }

}
