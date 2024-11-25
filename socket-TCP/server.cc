#include <SFML/Network.hpp>
#include <iostream>


int main(void){

    // server's listener
    sf::TcpListener listener;
    listener.listen(50011);

    // clients
    sf::TcpSocket socket;

    // connect server and client
    listener.accept(socket);

    // get client's IP address;
    std::cout << "New client connected: " << socket.getRemoteAddress() << std::endl;

    // Receive the message from the client;
    char buff[1024];
    std::size_t received = 0; // the length of the recevied message;
    socket.receive(buff, sizeof(buff) - 1, received);
    //std::cout << buff << "/" << sizeof(buff) << "/" << received << std::endl;
    std::cout << buff << std::endl;

    // Send back to client
    std::string message = "Welcome, client with IP: " + socket.getRemoteAddress().toString() + "!\0";
    socket.send(message.c_str(), message.size() + 1);

    socket.disconnect();

    return 0;
}