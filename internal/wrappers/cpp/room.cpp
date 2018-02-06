#include "room.hpp"
#include "collection.hpp"

namespace kuzzleio {
    Room::Room(room* r) {
      this->_room = r;
    }

    Room::Room(Collection *collection, json_object *filters, room_options* options) {
      this->_room = new room();
      room_new_room(this->_room, collection->_collection, filters, options);
    }

    Room::~Room() {
      unregisterRoom(this->_room);
      delete(this->_room);
    }

    int Room::count() Kuz_Throw_KuzzleException {
      int_result *r = room_count(_room);
      if (r->error != NULL)
          throwExceptionFromStatus(r);
      int ret = r->result;
      delete(r);
      return ret;
    }
}