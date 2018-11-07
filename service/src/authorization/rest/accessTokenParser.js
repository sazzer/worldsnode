// @flow

import {
    type $Request,
    type $Response,
    type NextFunction,
} from 'express';
import pino from 'pino';
import {type AccessTokenDeserializer} from '../serializer';
import {type ProblemType} from '../../rest/problem';

/** The logger to use */
const logger = pino();

/** The prefix for Bearer tokens */
const BEARER_PREFIX = 'Bearer ';

/** Response value indicating there was an authentication problem */
const authenticationProblem : ProblemType = {
    type: 'tag:grahamcox.co.uk,2018,worlds/problems/authentication/invalid_bearer_token',
    title: 'Error validating bearer token',
    status: 401,
};

/**
 * The actual middleware for parsing an access token from the request
 *
 * @param {$Request} req The incoming request
 * @param {$Response} res The outgoing response
 * @param {$NextFunction} next The next part of the chain to call
 * @param {AccessTokenDeserializer} deserializer The means to deserialize an Access Token
 * @return {any} The response from the next() handler
 */
export async function accessTokenParser(req: $Request, res: $Response, next: NextFunction,
    deserializer: AccessTokenDeserializer): any {
    const authHeader = req.headers['authorization'];
    logger.debug({authHeader}, 'Parsing Auth Header');
    if (!authHeader) {
        // We've not got an Authorization value
        logger.debug('No Auth Header received');
        return next();
    }

    if (!authHeader.startsWith(BEARER_PREFIX)) {
        // The Authorization value is not a Bearer token
        logger.debug({authHeader}, 'No bearer token received');
        return next();
    }

    const token = authHeader.substring(BEARER_PREFIX.length);
    logger.debug({token}, 'Received bearer token');

    try {
        const decoded = await deserializer.deserialize(token);
        logger.debug({accessToken: decoded}, 'Received access token');
        res.locals.accessToken = decoded;
        return next();
    } catch (e) {
        logger.warn({e}, 'Failed to decode access token');

        return res
            .status(401)
            .type('application/problem+json')
            .send(authenticationProblem);
    }
}

/**
 * Build the middleware function to use
 *
 * @param {AccessTokenDeserializer} deserializer The means to deserialize an Access Token
 * @return {Function} the middleware
 */
export default function buildAccessTokenParser(deserializer: AccessTokenDeserializer) :
    ($Request, $Response, NextFunction) => any {
    return (req, res, next) => accessTokenParser(req, res, next, deserializer);
}
