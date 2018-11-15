// @flow

import {type Resource} from '../../model/resource';
import {type UserData} from '../user';

/** The type representing a login the user has with a provider */
export type UserLoginModel = {
    '@type': string,
    'provider': string,
    'providerId': string,
    'providerName': string
};

/** The type representing the API Model of a User */
export type UserModel = {
    '@context': string,
    '@id': string,
    '@type': Array<string>,
    'created': Date,
    'updated': Date,
    'name': string,
    'email': ?string,
    'logins': Array<UserLoginModel>
};

/**
 * Actually translate the provided user to the model representation
 * @param user The user to translate
 * @return the API Model representation
 */
export function translateUserToModel(user: Resource<UserData>): UserModel {
    return {
        '@context': '/jsonld/users/User.jsonld',
        '@id': `/api/users/${user.identity.id}`,
        '@type': [
            'Person',
            'User',
        ],

        'created': user.identity.created,
        'updated': user.identity.updated,

        'name': user.data.name,
        'email': user.data.email,
        'logins': user.data.logins.map((provider) => ({
            '@type': 'Login',
            'provider': provider.provider,
            'providerId': provider.providerId,
            'providerName': provider.displayName,
        })),
    };
}
