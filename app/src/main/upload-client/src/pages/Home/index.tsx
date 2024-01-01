import { useEffect, useMemo } from "preact/hooks";
import preactLogo from "../../assets/preact.svg";
import "./style.css";
import Uppy from "@uppy/core";
import DragDrop from "@uppy/drag-drop";
import Tus from "@uppy/tus";

import "@uppy/core/dist/style.min.css";
import "@uppy/drag-drop/dist/style.min.css";
import ProgressBar from "@uppy/progress-bar";

export function Home() {
  useEffect(() => {
    const uppy = new Uppy({
      debug: true,
      autoProceed: true,
      allowMultipleUploads: true,
    })
      .use(DragDrop, { target: "#drag-drop" })
      .use(ProgressBar, { target: "#progress" })
      .use(Tus, {
        endpoint: "/api/upload",
        onBeforeRequest: async (...args) => console.log(args),
      });
  }, []);

  return (
    <div class="home">
      <div id="drag-drop"></div>
      <div id="progress"></div>
    </div>
  );
}
