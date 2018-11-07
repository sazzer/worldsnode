import express from 'express';
import bodyParser from 'body-parser';
import compression from 'compression';
import connectRid from 'connect-rid';
import cors from 'cors';
import errorHandler from 'errorhandler';
import responseTime from 'response-time';
import helmet from 'helmet';
import printRoutes from 'express-routemap';
import pino from 'express-pino-logger';

import type {
    $Request,
    $Response,
    NextFunction,
    Router,
} from 'express';

/** Type to represent something that can register routes with the system */
export type HandlerRegistration = (Router) => void;

/**
 * Build the service that we are going to run
 * @param middleware Any middleware to register
 * @param handlers The handlers to use for the service
 * @return the actual webapp service
 */
export default function buildService(middleware: Array<($Request, $Response, NextFunction) => any>,
    handlers: Array<HandlerRegistration>) {
    const app = express();

    app.use(responseTime());
    app.use(bodyParser.urlencoded({extended: false}));
    app.use(bodyParser.json());
    app.use(compression());
    app.use(connectRid());
    app.use(cors());
    app.use(errorHandler());
    app.use(helmet());
    app.use(pino());
    middleware.forEach((mw) => app.use(mw));

    const router = new express.Router();
    handlers.forEach((handler) => handler(router));
    app.use(router);

    printRoutes(app);
    return app;
}
