// @flow

import {
    type $Request,
    type $Response,
    type Router
} from 'express';
import checkHealth from './handler';
import type { HealthChecker } from '../healthcheck';

/**
 * Build the route registration mechanism
 * @param healthchecks The set of healthchecks to use
 */
export default function buildRegisterRoutes(healthchecks: Array<HealthChecker>) {
    return function registerRoutes(router: Router) {
        router.route('/health')
            .get((req: $Request, res: $Response) => checkHealth(res, healthchecks));
    }
}
