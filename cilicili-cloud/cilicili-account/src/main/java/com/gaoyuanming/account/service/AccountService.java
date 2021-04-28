package com.gaoyuanming.account.service;

import com.gaoyuanming.account.model.entity.Account;

import java.io.IOException;

public interface AccountService {

    Account login(String code) throws IOException;

    Account update(Account account);

    Account findById(String id);

}
