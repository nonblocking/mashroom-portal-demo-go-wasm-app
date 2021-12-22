
const bootstrap = (portalAppHostElement, portalAppSetup, clientServices) => {
    const { resourcesBasePath, appConfig: { message, pingButtonLabel }} = portalAppSetup;
    const { messageBus } = clientServices;

    return new Promise((resolve, reject) => {
        if (WebAssembly && typeof WebAssembly.instantiate === 'function') {
            const appInstanceId = 'DemoGoApp_' + Math.trunc(Math.random() * 1000000);
            if (!portalAppHostElement.id) {
                portalAppHostElement.id = appInstanceId;
            }
            const hostElementId = portalAppHostElement.id;

            // Expose the message bus instance globally to allow access from Go
            const messageBusGlobalObjName = `__${appInstanceId}`;
            window[messageBusGlobalObjName] = messageBus;

            // 'Go' is exposed by go_wasm_exec.js
            const go = new Go();
            let exitCode = 0;
            go.argv = [hostElementId, resourcesBasePath, message, pingButtonLabel, messageBusGlobalObjName];
            go.exit = (code) => {
                exitCode = code;
            };

            WebAssembly.instantiateStreaming(fetch(resourcesBasePath + '/main.wasm'), go.importObject).then(
                (result) => {
                    go.run(result.instance);
                    if (exitCode) {
                        reject(new Error('Go app exited with code: ' +  exitCode))
                    } else {
                        resolve({
                            willBeRemoved: () => {
                                console.info('Ummounting Go WebAssembly app');
                                // Send stop signal to Go app via message bus
                                messageBus.publish(`stop_${hostElementId}`, {});
                            }
                        });
                    }
                }, (error) => {
                    reject(error);
                });
        } else {
            reject(new Error('WebAssembly not supported by the Browser!'));
        }
    });
};

global.startGoWasmDemoApp = bootstrap;
