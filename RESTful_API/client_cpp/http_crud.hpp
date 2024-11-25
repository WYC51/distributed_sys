#ifndef HTTP_CRUD_HPP
#define HTTP_CRUD_HPP

#include <string>
#include "cpp-httplib/httplib.h"
#include "json/single_include/nlohmann/json.hpp"
#include <iostream>

using json = nlohmann::json;

class RestClient {

    public:
        httplib::Client cli;
        RestClient(httplib::Client& cli) : cli(std::move(cli)) {
            std::cout << "Create RESTful API Client" << std::endl;
            this -> cli.set_keep_alive(false);
        }
        std::string get_data(httplib::Client& cli, int id);
        std::string post_data(httplib::Client& cli, json& data);
        std::string put_data(httplib::Client& cli, int id, json& data);
        std::string delete_data(httplib::Client& cli, int id);
    
};

#endif
