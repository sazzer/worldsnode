// @flow

import type {
    Router
} from 'express';
import checkHealth from './handler';

/**
 * Register all of the routes with the given router
 * @param {Router} router
 */
export function registerRoutes(router: Router) {
    router.route('/health')
        .get(checkHealth);
}
