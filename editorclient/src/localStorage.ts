export function loadValue(name: string): Record<string, unknown> | undefined {
  try {
    const serializedValue = localStorage.getItem(name);
    if (serializedValue === null) {
      return undefined;
    }
    return JSON.parse(serializedValue);
  } catch (error) {
    return undefined;
  }
}

export function loadState(): Record<string, unknown> | undefined {
  return loadValue('state');
}

export function saveValue(name: string, value: string): void {
  try {
    const serializedValue = JSON.stringify(value);
    localStorage.setItem(name, serializedValue);
  } catch (error) {
    // die
  }
}

export function saveState(state: string): void {
  saveValue('state', state);
}

function removeValue(name: string) {
  try {
    localStorage.removeItem(name);
  } catch (error) {
    // die
  }
}

export function removeState(): void {
  removeValue('state');
}
