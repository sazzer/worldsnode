// @flow

/** Representation of an Access Token */
export type AccessToken = {
    /** The ID of the token itself */
    tokenId: string,
    /** The ID of the user that the token is for */
    userId: string,
    /** When the token was created */
    created: Date,
    /** When the token expires */
    expires: Date
};
