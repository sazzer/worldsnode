// @flow

import {Request, Response} from 'jest-express';
import {accessTokenParser} from '../accessTokenParser';

describe('accessTokenParser', () => {
    const deserializer = {
        deserialize: jest.fn(),
    };

    test('With no Authorization header', async () => {
        const request = new Request('/api/debug/accessToken', {
            headers: {

            },
        });
        const response = new Response();
        const next = jest.fn();

        await accessTokenParser(request, response, next, deserializer);
        expect(next).toBeCalled();
    });
    test('With an Authorization header that isn\'t a Bearer token', async () => {
        const request = new Request('/api/debug/accessToken', {
            headers: {
                'authorization': 'Basic abc123',
            },
        });
        const response = new Response();
        const next = jest.fn();

        await accessTokenParser(request, response, next, deserializer);
        expect(next).toBeCalled();
    });
    test('With a Bearer Token that isn\'t valid', async () => {
        const request = new Request('/api/debug/accessToken', {
            headers: {
                'authorization': 'Bearer abc123',
            },
        });
        const response = new Response();
        const next = jest.fn();

        deserializer.deserialize.mockReturnValueOnce(Promise.reject(''));
        await accessTokenParser(request, response, next, deserializer);

        expect(deserializer.deserialize).toBeCalledWith('abc123');
        expect(next).not.toBeCalled();
        expect(response.status).toBeCalledWith(401);
        expect(response.type).toBeCalledWith('application/problem+json');
        expect(response.send).toBeCalledWith({
            type: 'tag:grahamcox.co.uk,2018,worlds/problems/authentication/invalid_bearer_token',
            title: 'Error validating bearer token',
            status: 401,
        });
    });
    test('With a Bearer Token that is valid', async () => {
        const request = new Request('/api/debug/accessToken', {
            headers: {
                'authorization': 'Bearer abc123',
            },
        });
        const response = new Response();
        const next = jest.fn();

        deserializer.deserialize.mockReturnValueOnce(Promise.resolve({
            userId: 'someUserId',
            tokenId: 'generated-uuid',
            created: new Date('1970-01-01T00:00:00.000Z'),
            expires: new Date('1971-01-01T00:00:00.000Z'),
        }));
        await accessTokenParser(request, response, next, deserializer);

        expect(deserializer.deserialize).toBeCalledWith('abc123');
        expect(next).toBeCalled();
        expect(response.status).not.toBeCalled();
        expect(response.type).not.toBeCalled();
        expect(response.send).not.toBeCalled();
    });
});
