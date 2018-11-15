// @flow

import {type Resource} from '../../model/resource';
import {type UserData} from '../user';
import {type Database} from '../../db';
import pino from 'pino';

const logger = pino();
export type UserRetriever = {
    getById(id: string): Promise<Resource<UserData>>,
};

/**
 * Get the details of a single User from the database
 * @param database The database connection to use
 * @param {string} id The ID of the user to retrieve
 * @return the user
 */
export function getById(database: Database, id: string): Promise<Resource<UserData>> {
    return database.query({
        sql: 'SELECT * FROM users WHERE user_id = $1',
        binds: [id],
    }).then((results) => {
        logger.info({results}, 'User found');
    }).then(() => {
        return {
            identity: {
                id: id,
                version: '1',
                created: new Date(),
                updated: new Date(),
            },
            data: {
                name: 'Graham',
                email: 'graham@grahamcox.co.uk',
                logins: [
                    {
                        provider: 'google',
                        providerId: '123456',
                        displayName: 'graham@grahamcox.co.uk',
                    },
                ],
            },
        };
    });
}

/**
 * Mechanism to retrieve users from the database
 * @param database The database connection to use
 * @return The user retriever
 */
export default function buildUserRetriever(database: Database): UserRetriever {
    return {
        getById: (id) => getById(database, id),
    };
}
