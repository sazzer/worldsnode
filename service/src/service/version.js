// @flow

import type {
    Router,
    $Request,
    $Response
} from 'express';

/**
 * Register all of the routes with the given router
 * @param {Router} router
 */
export function registerRoutes(router: Router) {
    router.route('/api/version')
        .get((req: $Request, res: $Response) => res.send({
            version: '1.0.0'
        }));
}
