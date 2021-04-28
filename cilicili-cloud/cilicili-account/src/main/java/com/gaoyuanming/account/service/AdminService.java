package com.gaoyuanming.account.service;

import com.gaoyuanming.account.model.entity.Admin;

import java.util.List;

public interface AdminService {

    Admin save(Admin admin);

    List<Admin> findAll();

    void delete(String id);

    Admin loginByPassword(Admin admin) throws Exception;

    Admin loginByCode(String email, String code) throws Exception;

}
