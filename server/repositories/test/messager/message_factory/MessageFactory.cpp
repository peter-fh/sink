#include "MessageFactory.h"
#include <iostream>
#include <ostream>

MessageFactory::MessageFactory(){}
MessageFactory::~MessageFactory(){
    for (auto m: this->messages){
	delete m;
    }
}

void MessageFactory::addMessage(Message* msg){
    messages.push_back(msg);
}


std::ostream& operator<<(std::ostream& o, MessageFactory& mf){
    for (auto m: mf.messages){
	o << *m << std::endl;
    }
    return o;
}
