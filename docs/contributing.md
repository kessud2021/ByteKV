# Contributing

We welcome contributions! You can help by reporting issues, adding features, or improving documentation.

## Python Client

```bash
git clone https://github.com/yourusername/bytekv.git
cd bytekv/pybytekv
python -m venv venv
source venv/bin/activate   # macOS/Linux
venv\Scripts\activate      # Windows
pip install -r requirements.txt
python -m unittest discover tests
```

## Node.js Client
```
cd bytekv/node-client
npm install
npm test
npm run build   # if using TypeScript or bundler
npm link        # for local testing
```

## Tips

- Follow semantic versioning.

- Update README and docs for new commands.

- Write unit tests for Python and Node.js clients.

- Ensure cross-language compatibility.