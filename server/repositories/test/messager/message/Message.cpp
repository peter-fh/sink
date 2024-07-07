#include "Message.h"

Message::Message(std::string message){
    this->message = new std::string(message);
}

Message::~Message(){
    delete message;
}

std::ostream& operator<<(std::ostream& o, Message m){
    o << *m.message;
    return o;
}
