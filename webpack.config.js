
module.exports = {
    entry: __dirname + '/src/js',
    output: {
        path: __dirname + '/public',
        filename: 'index.js',
    },
    bail: true,
    module: {
        rules: [
            {
                test: /\.(js)$/,
                exclude: /node_modules/,
                use: [
                    {
                        loader: 'babel-loader',
                    },
                ],
            },
        ],
    },
    externals: [],
    resolve: {
        extensions: ['.js'],
    },
    devServer: {
        host: '0.0.0.0',
        allowedHosts: 'all',
        port: 8055,
        static: 'public',
        open: true
    },
};
