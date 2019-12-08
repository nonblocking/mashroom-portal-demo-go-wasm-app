
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
        extensions: ['.js', '.go'],
    },
    devServer: {
        inline: true,
        host: '0.0.0.0',
        disableHostCheck: true,
        port: 8055,
        contentBase: 'public',
        open: true,
    },
};
