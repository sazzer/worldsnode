// @flow

import {Pool} from 'pg';
import pino from 'pino';

export type Query = {
    sql: string,
    binds?: Array<any>,
}

/** The result of a query */
export type QueryResult = Array<{ [string] : ?any }>;

export type Database = {
    query: (Query) => Promise<QueryResult>,
};

/** The logger to use */
const logger = pino();

/**
 * Execute a query against the database and return the results
 *
 * @export
 * @param {Pool} pool The connection pool to use
 * @param {Query} query The query to execute
 * @return The query results
 */
export function executeQuery(pool: Pool, query: Query): Promise<QueryResult> {
    logger.debug({...query}, 'Executing query');
    return pool.query(query.sql, query.binds)
        .then((res) => {
            logger.debug({...query, result: res}, 'Executed query');
            return res.rows;
        })
        .catch((e) => {
            logger.warn({...query, error: e}, 'Error executing query');
            throw e;
        });
}

/**
 * Create the database connection
 *
 * @export
 * @param {string} uri The connection string
 * @return {Database} The database wrapper
 */
export default function buildDatabase(uri: string): Database {
    logger.info({uri}, 'Creating database connection');
    const pool = new Pool({
        connectionString: uri,
    });

    return {
        query: (query: Query) => executeQuery(pool, query),
    };
}
