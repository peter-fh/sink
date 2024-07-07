#ifndef MESSAGE_H
#define MESSAGE_H
#include <ostream>
#include <string>
class Message{
public:
    Message(std::string message);
    ~Message();
    friend std::ostream& operator<<(std::ostream&,Message);
private:
    std::string *message;
};
#endif
