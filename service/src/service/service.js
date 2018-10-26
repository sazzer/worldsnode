import express from 'express';
import bodyParser from 'body-parser';
import compression from 'compression';
import connectRid from 'connect-rid';
import cors from 'cors';
import errorHandler from 'errorhandler';
import responseTime from 'response-time';
import helmet from 'helmet';
import printRoutes from 'express-routemap';
import morgan from 'morgan';

import type {
    Router
} from 'express';

/** Type to represent something that can register routes with the system */
export type HandlerRegistration = (Router) => void;

/**
 * Build the service that we are going to run
 */
export default function buildService(handlers: Array<HandlerRegistration>) {
    const app = express();

    app.use(responseTime());
    app.use(bodyParser.urlencoded({ extended: false }));
    app.use(bodyParser.json());
    app.use(compression());
    app.use(connectRid());
    app.use(cors());
    app.use(errorHandler());
    app.use(helmet());
    app.use(morgan('dev'));

    const router = express.Router();
    handlers.forEach(handler => handler(router));
    app.use(router);

    printRoutes(app);
    return app;
}
