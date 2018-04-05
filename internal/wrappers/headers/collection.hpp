#ifndef _COLLECTION_HPP_
#define _COLLECTION_HPP_

#include <iostream>
#include <list>
#include "core.hpp"

namespace kuzzleio {

    class Collection {
        collection* _collection;
        Collection();

        public:
            Collection(Kuzzle* kuzzle);
            Collection(Kuzzle* kuzzle, collection *collection);
            virtual ~Collection();
            void create(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            bool exists(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string list(const std::string& index, query_options *options=NULL) Kuz_Throw_KuzzleException;
            void truncate(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string getMapping(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            void updateMapping(const std::string& index, const std::string& collection, const std::string& body, query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string getSpecifications(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
            search_result* searchSpecifications(query_options *options=NULL) Kuz_Throw_KuzzleException;
            std::string updateSpecifications(const std::string& index, const std::string& collection, const std::string& body, query_options *options=NULL) Kuz_Throw_KuzzleException;
            bool validateSpecifications(const std::string& body, query_options *options=NULL) Kuz_Throw_KuzzleException;
            void deleteSpecifications(const std::string& index, const std::string& collection, query_options *options=NULL) Kuz_Throw_KuzzleException;
    };
}

#endif
