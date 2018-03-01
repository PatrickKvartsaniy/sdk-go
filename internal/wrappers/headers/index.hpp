#ifndef _KUZZLE_INDEX_HPP
#define _KUZZLE_INDEX_HPP

#include "exceptions.hpp"
#include "core.hpp"

namespace kuzzleio {
  class Index {
    index *_index;
    Index();

    public:
      Index(Kuzzle* kuzzle);
      virtual ~Index();
      void create(const std::string& index) Kuz_Throw_KuzzleException;
      void delete_(const std::string& index) Kuz_Throw_KuzzleException;
      string* mDelete(const std::string* indexes) Kuz_Throw_KuzzleException;
      bool exists(const std::string& index) Kuz_Throw_KuzzleException;
      void refresh(const std::string& index) Kuz_Throw_KuzzleException;
      void refreshInternal() Kuz_Throw_KuzzleException;
      void setAutoRefresh(const std::string& index, bool autoRefresh) Kuz_Throw_KuzzleException;
  };
}

#endif
