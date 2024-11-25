#include <iostream>
#include <string>
#include "cpp-httplib/httplib.h"
#include "http_crud.hpp"

std::string RestClient::get_data(httplib::Client& cli, int id = -1){
    std::string path = "/show_demo_data";
    if(id != -1){
        path += "/" + std::to_string(id);
        auto res = cli.Get(path.c_str());
        if(res){
            if(res->status == 200){
                return res->body;
            }
        }
    } else{
        auto res = cli.Get(path);
        if(res){
            if(res->status == 200){
                return res->body;
            }
        }
    }
    return "Error";
}

std::string RestClient::put_data(httplib::Client& cli, int id, json& data){
    std::string path = "/show_demo_data/" + std::to_string(id);
    auto headers = httplib::Headers{{"Content-Type", "application/json"}};
    auto res = cli.Put(path, headers, data.dump(), "application/json");
    if (res) {
        if (res->status == 200) {
            return res->body;
        }
    }
    return "Error";

}

std::string  RestClient::post_data(httplib::Client& cli, json& data){
    auto headers = httplib::Headers{{"Content-Type", "application/json"}};
    auto res = cli.Post("/show_demo_data", headers, data.dump(), "application/json");
    if (res) {
        if (res->status == 200) {
            return res->body;
        }
    }
    return "Error";
}

std::string  RestClient::delete_data(httplib::Client& cli, int id){
    std::string path = "/show_demo_data/" + std::to_string(id);
    auto res = cli.Delete(path.c_str());
    if (res) {
        if (res->status == 200) {
            return res->body;
        }
    }
    return "Error";
}



int main(int argc, char* argv[]){
    httplib::Client cli("http://localhost:8080");
    RestClient client(cli);
    if (argc == 1) {
        auto res = client.cli.Get("/");
        if(res){
            if(res->status == 200){
                std::cout << res->body << std::endl;
            }
        }
        return 1;
    }
    else{
        if ((std::string(argv[1]) == "GET") && argc == 2){
            std::string data = client.get_data(client.cli);
            std::cout << data << std::endl;
        }
        else if ((std::string(argv[1]) == "GET") && argc == 3){
            std::string data = client.get_data(client.cli, std::stoi(argv[2]));
            std::cout << data << std::endl;
        }
        else if ((std::string(argv[1]) == "POST") && argc == 2){
            json ex1 = json::parse(R"(
            {
                "name": "test500",
                "id": 500,
                "mail": "500"
            }
            )");
            std::string data = client.post_data(client.cli, ex1);
            std::cout << data << std::endl;
        }
        else if ((std::string(argv[1]) == "PUT") && argc == 3){
            json ex1 = json::parse(R"(
            {
                "name": "test500",
                "id": 500,
                "mail": "500@update.com"
            }
            )");
            std::string data = client.put_data(client.cli, std::stoi(argv[2]), ex1);
            std::cout << data << std::endl;
        }
        else if ((std::string(argv[1]) == "DELETE") && argc == 3){
            std::string data = client.delete_data(client.cli, std::stoi(argv[2]));
            std::cout << data << std::endl;
        }
        else{
            std::cout << "Invalid arguments" << std::endl;
        }
    }

    return 0;
}