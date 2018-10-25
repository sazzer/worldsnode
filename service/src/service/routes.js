// @flow

import type { Router } from 'express';
import { registerRoutes as version } from './version';
import { registerRoutes as health } from '../health/rest';

/**
 * Register all of the routes with the given router
 * @param {Router} router
 */
export function registerRoutes(router: Router) {
    version(router);
    health(router);
}
