#include <SFML/Network.hpp>
#include <iostream>

int main(void){

    // connect with the server;
    sf::TcpSocket socket;
    sf::Socket::Status status = socket.connect("127.0.0.1", 50011);

    // send message to the server;
    std::string message = "Hello server, this message is from client";
    socket.send(message.c_str(), message.size() + 1);


    // receive the message from the server;
    char buff[1024];
    std::size_t received = 0; // the length of the recevied message;
    socket.receive(buff, sizeof(buff) - 1, received);
    buff[received] = '\0';
    //std::cout << buff << "/" << sizeof(buff) << "/" << received << std::endl;
    std::cout << buff << std::endl;

    socket.disconnect();
    return 0;
}