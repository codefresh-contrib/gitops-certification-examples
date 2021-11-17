package com.example.springboot;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;

@RestController
public class HelloController {

	String my_sql_db = System.getenv("MY_MYSQL_DB");

	@RequestMapping("/")
	public String index() {
		return "Mysql is at " +my_sql_db;
	}

}
