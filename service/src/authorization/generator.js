// @flow

import uuid from 'uuid/v4';
import {DateTime, Duration} from 'luxon';
import {type AccessToken} from './accessToken';

/** The type reprsenting the Access Token Generator */
export type AccessTokenGenerator = {
    generate: (string) => AccessToken,
};

/**
 * Generate an access token for the given User ID
 *
 * @param {string} userId The ID of the user the token is for
 * @param {Duration} tokenDuration The duration of the token
 * @return {AccessToken} the token
 */
export function generate(userId: string, tokenDuration: Duration): AccessToken {
    const now = DateTime.utc();
    const expires = now.plus(tokenDuration);

    return {
        tokenId: uuid(),
        userId: userId,
        created: now.toJSDate(),
        expires: expires.toJSDate(),
    };
}

/**
 * Build the Access Token Generator to use
 *
 * @export
 * @param {string} duration the duration of the access tokens
 * @return {AccessTokenGenerator} the Access Token Generator
 */
export default function buildAccessTokenGenerator(duration: string): AccessTokenGenerator {
    const tokenDuration = Duration.fromISO(duration);
    return {
        generate: (userId) => generate(userId, tokenDuration),
    };
}
