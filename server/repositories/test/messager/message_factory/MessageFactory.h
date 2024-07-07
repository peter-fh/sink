#include <ostream>
#include <string>
#include <vector>
#include "../message/Message.h"


class MessageFactory {
public:
    MessageFactory();
    ~MessageFactory();
    void addMessage(Message* msg);
    friend std::ostream& operator<<(std::ostream&, MessageFactory&);
private:

    std::vector<Message*> messages;

};
