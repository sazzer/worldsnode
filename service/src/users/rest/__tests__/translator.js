// @flow

import * as testSubject from '../translator';

describe('translateUserToModel', () => {
    test('Translating a minimal user', () => {
        const user = {
            identity: {
                id: '123',
                version: '1',
                created: new Date('2018-11-15T09:34:00Z'),
                updated: new Date('2018-11-15T09:34:00Z'),
            },
            data: {
                name: 'Graham',
                logins: [],
            },
        };
        const translated = testSubject.translateUserToModel(user);
        expect(translated).toMatchSnapshot();
    });

    test('Translating a full user', () => {
        const user = {
            identity: {
                id: '123',
                version: '1',
                created: new Date('2018-11-15T09:34:00Z'),
                updated: new Date('2018-11-15T09:34:00Z'),
            },
            data: {
                name: 'Graham',
                email: 'graham@grahamcox.co.uk',
                logins: [
                    {
                        provider: 'google',
                        providerId: '123456',
                        displayName: 'graham@grahamcox.co.uk',
                    },
                ],
            },
        };
        const translated = testSubject.translateUserToModel(user);
        expect(translated).toMatchSnapshot();
    });
});
