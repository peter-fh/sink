#include "messager/message_factory/MessageFactory.h"
#include <iostream>

int main(){
    MessageFactory* message_factory = new MessageFactory;
    Message* hello_world = new Message("Hello World!");
    message_factory->addMessage(hello_world);
    std::cout << *message_factory;
}
