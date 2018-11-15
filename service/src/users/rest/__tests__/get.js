// @flow

import {Request, Response} from 'jest-express';
import * as testSubject from '../get';

describe('GET /api/users/123', () => {
    test('When the user exists', async () => {
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

        const request = new Request('/api/users/123');
        request.setParams({id: '123'});
        const response = new Response();
        const retriever = {
            getById: jest.fn(() => Promise.resolve(user)),
        };
        await testSubject.default(request, response, retriever);

        expect(response.type).toBeCalledWith('application/json');
        expect(response.status).toBeCalledWith(200);
        expect(response.send).toMatchSnapshot();
        expect(retriever.getById).toBeCalledWith('123');
    });
    test('When the user doesn\'t exist', () => {

    });
    test('When the ID is malformed', () => {

    });
    test('When the database is missing', () => {

    });
});
