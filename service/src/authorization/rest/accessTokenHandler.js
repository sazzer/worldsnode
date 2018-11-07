// @flow

import {
    type $Request,
    type $Response,
    type Router,
} from 'express';
import pino from 'pino';
import {type AccessTokenSerializer} from '../serializer';
import {type AccessTokenGenerator} from '../generator';

/** The logger to use */
const logger = pino();

/** Type to represent the request body for requesting an access token */
type AccessTokenRequestBody = {
    userId: string
};

/**
 * Handler that will build an access token for the requested User ID
 *
 * @export
 * @param {$Request} req The incoming request
 * @param {$Response} res The outgoing response
 * @param {AccessTokenSerializer} serializer The means to serialize an access token
 * @param {AccessTokenGenerator} generator The means to generate an access token
 */
export async function buildAccessToken(req: $Request, res: $Response,
    serializer: AccessTokenSerializer, generator: AccessTokenGenerator) {
    const body = ((req.body: any): AccessTokenRequestBody);
    const user: string = body.userId;
    const token = generator.generate(user);
    const serialized = await serializer.serialize(token);

    logger.warn({user, token, serialized}, 'Building Access Token for User');
    res.send({
        token: serialized,
    });
}

/**
 * Build the means to register the route for getting an access token
 *
 * @export
 * @param {AccessTokenSerializer} serializer The means to serialize an access token
 * @param {AccessTokenGenerator} generator The means to generate an access token
 * @return {function} Function for registering routes
 */
export default function buildRegisterRoutes(serializer: AccessTokenSerializer, generator: AccessTokenGenerator) {
    return function registerRoutes(router: Router) {
        router.route('/api/debug/accessToken')
            .post((req: $Request, res: $Response) => buildAccessToken(req, res, serializer, generator));
    };
}
