import type { Frame, PathItem } from "../common/types";

export function configurePath(
  firstFrame: Pick<Frame, "uid">,
  frameByUid: Record<string, Frame>,
  branchNumByUid?: Record<string, number>
): PathItem[] {
  let frameUid = firstFrame.uid;
  const path: PathItem[] = [];
  while (path.length <= Object.keys(frameByUid).length) {
    const pathItem: PathItem = {
      frameUid,
      branchNum: branchNumByUid?.[frameUid] || 0,
    };
    path.push(pathItem);

    const actions = frameByUid[frameUid]?.actions;
    if (actions == null || !actions.length) break;

    const nextFrame = actions[pathItem.branchNum].nextFrame;
    if (!nextFrame || !frameByUid[nextFrame.uid]) break;

    frameUid = nextFrame.uid;
  }

  return path;
}
