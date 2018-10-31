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

/** Type representing a means to serialize an access token */
export type AccessTokenSerializer = {
    serialize: (AccessToken) => Promise<string>,
};

/** Type representing a means to deserialize an access token */
export type AccessTokenDeserializer = {
    deserialize: (string) => Promise<AccessToken>,
};

/**
 * Serialize the given Access Token to a String
 *
 * @param {AccessToken} accessToken The access token to serialize
 * @param {string} signingKey The signing key to use
 * @return {Promise<string>} the Access Token as a string
 * @memberof AccessTokenSerializer
 */
export function serialize(accessToken: AccessToken, signingKey: string): Promise<string> {
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
        signingKey,
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
 * @param {string} signingKey The signing key to use
 * @return {Promise<AccessToken>} the Access Token
 * @memberof AccessTokenSerializer
 */
export function deserialize(token: string, signingKey: string): Promise<AccessToken> {
    logger.debug({token}, 'Deserializing access token');
    return new Promise((resolve, reject) => {
        jwt.verify(token,
            signingKey,
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

/**
 * Actually build the Access Token Serializer to use with the given key
 *
 * @export
 * @param {string} key the key to use
 * @return {AccessTokenSerializer} the Access Token Serializer to use
 */
export default function buildAccessTokenSerializer(key: string): AccessTokenSerializer & AccessTokenDeserializer {
    return {
        serialize: (accessToken) => serialize(accessToken, key),
        deserialize: (token) => deserialize(token, key),
    };
}
