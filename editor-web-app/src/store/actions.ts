import type { State } from "./state";
import type {
  Script,
  ScriptInfo,
  TraversableScript,
  Frame,
} from "@/common/types";
import {
  getScriptsInfo,
  getScript,
  updateScript,
  deleteScript,
} from "@/common/requests";

type Actions = {
  setScriptsInfo(this: State & Actions, scriptsInfo: ScriptInfo[]): void;
  setScript(this: State & Actions, script: TraversableScript): void;
  updateScript(
    this: State & Actions,
    data: {
      script?: Partial<Script>;
      frames?: Frame[];
      uids?: Record<string, string> | null;
      frameIdsToDel?: string[];
      actionIdsToDel?: string[];
    }
  ): Promise<void>;
  selectFrame(this: State & Actions, uid?: string): void;
  configurePath(
    this: State & Actions,
    fork: {
      frameUid: string;
      branchNum: number;
    }
  ): void;
  loadScriptsInfo(this: State & Actions): Promise<void>;
  loadScript(this: State & Actions, uid: string): Promise<void>;
  updateScript(
    this: State & Actions,
    data: {
      script?: Script;
      frames?: Frame[];
      uids?: Record<string, string>;
      frameIdsToDel?: string[];
      actionIdsToDel?: string[];
    }
  ): Promise<void>;
  deleteScript(this: State & Actions, uid: string): Promise<void>;
};

export const actions: Actions = {
  setScriptsInfo(scriptsInfo) {
    this.scriptsInfo = scriptsInfo;
  },
  setScript(script) {
    this.script = script;
  },
  selectFrame(uid) {
    this.frameUid = uid;
  },
  configurePath(fork) {
    this.script.branchNumByUid[fork.frameUid] = fork.branchNum;
  },
  loadScriptsInfo() {
    return new Promise((resolve, reject) => {
      getScriptsInfo()
        .then((data) => {
          this.setScriptsInfo(data.scripts);
          resolve();
        })
        .catch(reject);
    });
  },
  loadScript(uid) {
    return new Promise((resolve, reject) => {
      getScript(uid)
        .then((data) => {
          const script: Script = data.script;

          const frameByUid: Record<string, Frame> = {};
          script.frames.forEach((frame) => {
            frameByUid[frame.uid] = frame;
          });

          const traversableScript: TraversableScript = {
            ...script,
            frameByUid,
            path: [],
            branchNumByUid: {},
          };
          this.setScript(traversableScript);
          this.selectFrame(traversableScript.firstFrame.uid);
          resolve();
        })
        .catch((err) => {
          reject(err);
        });
    });
  },
  updateScript({ script, frames, frameIdsToDel, actionIdsToDel }) {
    return new Promise((resolve, reject) => {
      if (!this.script.uid) return;

      return updateScript(
        { uid: this.script.uid, ...script, frames } as Script,
        { frameIdsToDel, actionIdsToDel }
      )
        .then((res) => {
          this.script = {
            ...this.script,
            ...script,
          };

          const framesWithCorrectUids = !res.uids
            ? frames
            : frames?.map((frame) => {
                const actionsWithCorrectUids = frame.actions?.map((action) => {
                  const newNextFrameUid = action.nextFrame
                    ? res.uids?.[action.nextFrame?.uid.slice(2)]
                    : null;
                  const nextFrameWithCorrectUid = newNextFrameUid
                    ? { uid: newNextFrameUid }
                    : action.nextFrame;

                  const newActionUid = res.uids?.[action.uid.slice(2)];
                  return newActionUid
                    ? {
                        ...action,
                        uid: newActionUid,
                        nextFrame: nextFrameWithCorrectUid,
                      }
                    : { ...action, nextFrame: nextFrameWithCorrectUid };
                });

                const newFrameUid = res.uids?.[frame.uid.slice(2)];
                return newFrameUid
                  ? {
                      ...frame,
                      uid: newFrameUid,
                      actions: actionsWithCorrectUids,
                    }
                  : { ...frame, actions: actionsWithCorrectUids };
              });

          frameIdsToDel?.forEach((frameId) => {
            delete this.script.frameByUid[frameId];
          });

          Object.values(this.script.frameByUid).forEach((frame: Frame) => {
            if (frame.actions) {
              frame["actions"] = frame.actions?.filter(
                (action) => !actionIdsToDel?.includes(action.uid)
              );
            }
          });

          framesWithCorrectUids?.forEach((frame) => {
            const currentFrame = this.script?.frameByUid[frame.uid];
            if (currentFrame) {
              const currentActions =
                currentFrame.actions?.map((currentAction) => {
                  const newAppropriateAction = frame.actions?.find(
                    (action) => action.uid === currentAction.uid
                  );
                  return { ...currentAction, ...newAppropriateAction };
                }) || [];

              const newActions =
                frame.actions?.filter(
                  (action) =>
                    !currentActions.find(
                      (currentAction) => currentAction.uid === action.uid
                    )
                ) || [];

              this.script.frameByUid[frame.uid] = {
                ...currentFrame,
                ...frame,
                actions: [...currentActions, ...newActions],
              };
            } else {
              this.script.frameByUid[frame.uid] = frame;
            }
          });
          resolve();
        })
        .catch((err) => {
          reject(err);
        });
    });
  },
  deleteScript(uid) {
    return new Promise((resolve, reject) => {
      return deleteScript(uid)
        .then(() => {
          this.scriptsInfo = this.scriptsInfo.filter(
            (script) => script.uid !== uid
          );
          resolve();
        })
        .catch((err) => {
          reject(err);
        });
    });
  },
};
