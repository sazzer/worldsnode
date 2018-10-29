// @flow

import uuid from 'uuid/v4';
import {DateTime, Duration} from 'luxon';
import {type AccessToken} from './accessToken';

/**
 * Mechanism to generate an access token for a user
 *
 * @export
 * @class AccessTokenGenerator
 */
export default class AccessTokenGenerator {
    /** The duration of the access token */
    _tokenDuration: Duration;

    /**
     * Construct the Access Token Generator
     * @param {string} expiry The duration of the access tokens
     */
    constructor(expiry: string) {
        this._tokenDuration = Duration.fromISO(expiry);
    }

    /**
     * Generate an access token for the given User ID
     *
     * @param {string} userId The ID of the user the token is for
     * @return {AccessToken} the token
     */
    generate(userId: string): AccessToken {
        const now = DateTime.utc();
        const expires = now.plus(this._tokenDuration);

        return {
            tokenId: uuid(),
            userId: userId,
            created: now.toJSDate(),
            expires: expires.toJSDate(),
        };
    }
}
