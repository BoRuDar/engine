/**
 * Created by igor on 07.03.17.
 */

"use strict";

var conf = require(__appRoot + '/conf'),
    collectionAgent= conf.get('mongodb:collectionDialerAgents');

module.exports = {
    addQuery: addQuery
};

function addQuery(db) {
    return {
        list: (domain, filter = {}, project = {}, cb) => {
            // TODO 
            // if (domain)
            //     filter.domain = domain;

            return db
                .collection(collectionAgent)
                .find(filter, project)
                .toArray(cb);
        }
    }
}