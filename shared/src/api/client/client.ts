import { NextObserver, Observable, Subscribable, Subscription, Unsubscribable } from 'rxjs'
import { switchMap } from 'rxjs/operators'
import { Connection } from '../protocol/jsonrpc2/connection'
import { createExtensionHostClientConnection, ExtensionHostClientConnection } from './connection'
import { Environment } from './environment'
import { Services } from './services'

export interface ExtensionHostClient extends Unsubscribable {
    /**
     * Closes the connection to the extension host and stops the controller from reestablishing new
     * connections.
     */
    unsubscribe(): void
}

/**
 * Creates a client to communicate with an extension host.
 *
 * TODO!2(sqs), make this for all shared code and the internal "extension" API, not just cross-context extensions.
 *
 * @param extensionHostConnection An observable that emits the connection to the extension host each time a new
 * connection is established.
 */
export function createExtensionHostClient(
    // TODO!(sqs): make it possible to just use an observable of environment, not
    // behaviorsubject, to simplify data flow
    environment: Subscribable<Environment> & { value: Environment } & NextObserver<Environment>,
    services: Services,
    extensionHostConnection: Observable<Connection>
): ExtensionHostClient {
    const subscriptions = new Subscription()

    const client = extensionHostConnection.pipe(
        switchMap(connection => {
            const client = createExtensionHostClientConnection(connection, environment, services)
            return new Observable<ExtensionHostClientConnection>(sub => {
                sub.next(client)
                return () => client.unsubscribe()
            })
        })
    )
    subscriptions.add(client.subscribe())

    return {
        unsubscribe: () => subscriptions.unsubscribe(),
    }
}
