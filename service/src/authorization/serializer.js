// @flow

import jwt from 'jsonwebtoken';
import pino from 'pino';
import {type AccessToken} from './accessToken';

/** The logger to use */
const logger = pino();

/** Type representing an error from serializing or deserializing a token */
export type SerializerError = {
    message: string
};

/**
 * Mechanism to serialize an access token to a string and back
 *
 * @export
 * @class AccessTokenSerializer
 */
export default class AccessTokenSerializer {
    /** The key used for signing the JWT */
    _signingKey: string;

    /**
     * Construct the serializer
     * @param {string} key The signing key to use
     */
    constructor(key: string) {
        this._signingKey = key;
    }

    /**
     * Serialize the given Access Token to a String
     *
     * @param {AccessToken} accessToken The access token to serialize
     * @return {Promise<string>} the Access Token as a string
     * @memberof AccessTokenSerializer
     */
    serialize(accessToken: AccessToken): Promise<string> {
        logger.debug({accessToken}, 'Serializing access token');
        return new Promise((resolve, reject) => {
            jwt.sign({
                aud: 'worlds',
                iss: 'worlds',
                jwtid: accessToken.tokenId,
                sub: accessToken.userId,
                iat: Math.floor(accessToken.created / 1000),
                exp: Math.floor(accessToken.expires / 1000),
            },
            this._signingKey,
            {
                algorithm: 'HS256',
            },
            (err, token) => {
                if (err) {
                    logger.warn({accessToken, err}, 'Failed to serialize access token');
                    reject({
                        message: err.message,
                    });
                } else {
                    logger.debug({accessToken, token}, 'Serialized access token');
                    resolve(token);
                }
            });
        });
    }

    /**
     * Deserialize the given String to an Access Token
     *
     * @param {string} token The token to deserialize
     * @return {Promise<AccessToken>} the Access Token
     * @memberof AccessTokenSerializer
     */
    deserialize(token: string): Promise<AccessToken> {
        logger.debug({token}, 'Deserializing access token');
        return new Promise((resolve, reject) => {
            jwt.verify(token,
                this._signingKey,
                {
                    algorithms: ['HS256'],
                    audience: 'worlds',
                    issuer: 'worlds',
                },
                (err, decoded) => {
                    if (err) {
                        logger.warn({token, err}, 'Failed to deserialize access token');
                        reject({
                            message: err.message,
                        });
                    } else {
                        logger.debug({token, decoded}, 'Deserialized access token');
                        resolve(decoded);
                    }
                });
        })
            .then((jwtToken) => (
                {
                    userId: jwtToken.sub,
                    tokenId: jwtToken.jwtid,
                    expires: new Date(jwtToken.exp * 1000),
                    created: new Date(jwtToken.iat * 1000),
                }
            ));
    }
}
