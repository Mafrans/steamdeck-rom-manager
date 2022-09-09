import { useState } from "preact/hooks";
import { Upload } from "tus-js-client";

type Progress = {
  current: number;
  target: number;
  completion: number;
};

export const useUpload = () => {
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState<Error>();
  const [progress, setProgress] = useState<Progress>({
    current: 0,
    target: 0,
    completion: 0,
  });

  return {
    success,
    error,
    progress,
    upload(file: File, name: string = file.name) {
      const upload = new Upload(file, {
        endpoint: `${location.origin}/files`,
        retryDelays: [0, 3000, 5000, 10000, 20000],
        metadata: {
          filename: name,
          filetype: file.type,
        },
        onSuccess() {
          setSuccess(true);
        },
        onError(error) {
          setError(error);
        },
        onProgress(current, target) {
          setProgress({
            current,
            target,
            completion: current / target,
          });
        },
      });

      upload.start();
    },
  };
};
