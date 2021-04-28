package com.gaoyuanming.account.controller;

import com.gaoyuanming.account.model.entity.Admin;
import com.gaoyuanming.account.service.AdminService;
import org.springframework.web.bind.annotation.*;

import javax.annotation.Resource;
import java.util.List;

@RestController
public class AdminController {

    @Resource
    private AdminService adminService;

    @PostMapping("/admin")
    public Admin save(@RequestBody Admin admin) {
        return adminService.save(admin);
    }

    @DeleteMapping("/admin/{id}")
    public void delete(@PathVariable String id) {
        adminService.delete(id);
    }

    @GetMapping("/admin")
    public List<Admin> findAll() {
        return adminService.findAll();
    }

    @PostMapping("/admin/loginByPassword")
    public Admin loginByPassword(@RequestBody Admin admin) throws Exception {
        return adminService.loginByPassword(admin);
    }

    @GetMapping("/admin/loginByEmail")
    public Admin loginByEmail(String email, String code) throws Exception {
        return adminService.loginByCode(email, code);
    }

}
