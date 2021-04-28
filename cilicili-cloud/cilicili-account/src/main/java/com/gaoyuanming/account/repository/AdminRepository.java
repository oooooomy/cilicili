package com.gaoyuanming.account.repository;

import com.gaoyuanming.account.model.entity.Admin;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

@Repository
public interface AdminRepository extends JpaRepository<Admin, String> {

    Admin findByEmail(String email);

    Admin findAllByEmailAndPassword(String email, String password);

}
