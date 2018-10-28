// @flow

import uuid from 'uuid/v4';
import {DateTime, Duration} from 'luxon';
import {type AccessToken} from './accessToken';

/**
 * Generate an access token for the given User ID
 *
 * @param {string} userId The ID of the user the token is for
 * @return {AccessToken} the token
 */
export default function generateAccessToken(userId: string): AccessToken {
    const tokenDuration = Duration.fromISO('P1Y');
    const now = DateTime.utc();
    const expires = now.plus(tokenDuration);

    return {
        tokenId: uuid(),
        userId: userId,
        created: now.toJSDate(),
        expires: expires.toJSDate(),
    };
}
