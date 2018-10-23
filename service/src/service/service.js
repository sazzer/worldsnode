import express from 'express';
import bodyParser from 'body-parser';
import compression from 'compression';
import connectRid from 'connect-rid';
import cors from 'cors';
import errorHandler from 'errorhandler';
import responseTime from 'response-time';
import helmet from 'helmet';
import printRoutes from 'express-routemap';
import { registerRoutes } from './routes';

/**
 * Build the service that we are going to run
 */
export function buildService() {
    const app = express();

    app.use(responseTime());
    app.use(bodyParser.urlencoded({ extended: false }));
    app.use(bodyParser.json());
    app.use(compression());
    app.use(connectRid());
    app.use(cors());
    app.use(errorHandler());
    app.use(helmet());

    const router = express.Router();
    registerRoutes(router);
    app.use(router);

    printRoutes(app);
    return app;
}
