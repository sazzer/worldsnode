// @flow

/** Type representing a login that a user has with a provider */
export type UserLogin = {
    /** The name of the provider */
    provider: string,
    /** The ID of the user at the provider */
    providerId: string,
    /** The display name of the user at the provider */
    displayName: string
};

/**
 * Type representing the data of a user
 */
export type UserData = {
    /** The name of the user */
    name: string,
    /** The email address of the user */
    email?: string,
    /** The logins that the user has */
    logins: Array<UserLogin>
};
