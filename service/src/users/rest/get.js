// @flow

import {
    type $Request,
    type $Response,
} from 'express';
import {type UserRetriever} from '../service/retriever';
import {translateUserToModel} from './translator';
import pino from 'pino';

/** The logger to use */
const logger = pino();
/**
 * Get the details of a single user from the database
 *
 * @export
 * @param {$Request} req the request
 * @param {$Response} res the response
 * @param {UserRetriever} retriever The user retriever
 */
export default async function getUser(req: $Request, res: $Response, retriever: UserRetriever) {
    const userId = req.params.id;

    const user = await retriever.getById(userId);
    const userModel = translateUserToModel(user);
    logger.info({userModel}, 'Retrieved user');

    res
        .type('application/json')
        .status(200)
        .send(userModel);
}
