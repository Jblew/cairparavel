import { interpret } from 'xstate';
import { CreateEventInterpreter, createEventMachine } from './machine';
import { saveEvent } from './services';

export function interpretCreateEventMachine(): CreateEventInterpreter {
  return interpret(createEventMachine.withConfig({
    services: {
      saveEvent: (ctx) => saveEvent(ctx.event)
    }
  }))
}
