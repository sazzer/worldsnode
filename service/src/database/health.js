// @flow

import pino from 'pino';
import {type Database} from './db';
import {HEALTH_PASS, HEALTH_FAIL, type Health, type HealthChecker} from '../health/healthcheck';

/** The logger to use */
const logger = pino();

/**
 * Healthcheck for the database
 *
 * @export
 * @param {Database} database The database connection
 * @return {Array<Health>} the healthcheck results
 */
export function databaseHealthcheck(database: Database): Array<Promise<Health>> {
    const databaseHealth = database.query({sql: 'SELECT 1'})
        .then(() => ({
            component: 'database',
            measurement: 'ping',
            status: HEALTH_PASS,
        }))
        .catch((e) => {
            logger.error({e}, 'Database healthcheck failed');
            return {
                component: 'database',
                measurement: 'ping',
                status: HEALTH_FAIL,
            };
        });

    return [databaseHealth];
}

/**
 * Build the database healthchecks to use
 *
 * @export
 * @param {Database} database The database connection
 * @return {HealthChecker} The healthchecker
 */
export default function buildDatabaseHealthcheck(database: Database): HealthChecker {
    return () => databaseHealthcheck(database);
}
