import { buildService } from './service';

const port = 3000;

const app = buildService();
app.listen(port, () => console.log(`Example app listening on port ${port}!`));
