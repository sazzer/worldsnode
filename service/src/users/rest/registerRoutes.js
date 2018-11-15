// @flow

import {
    type $Request,
    type $Response,
    type Router,
} from 'express';

import getUser from './get';
import {type UserRetriever} from '../service/retriever';

/**
 * Build the route registration mechanism
 * @param {UserRetriever} userRetriever the means to load user records
 * @return The mechanismm for registering routes
 */
export default function buildRegisterRoutes(userRetriever: UserRetriever) {
    return function registerRoutes(router: Router) {
        router.route('/api/users/:id')
            .get((req: $Request, res: $Response) => getUser(req, res, userRetriever));
    };
}
