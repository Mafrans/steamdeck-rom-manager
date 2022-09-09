import { useMemo, useState } from "preact/hooks";
import { useUpload } from "../../hooks/useUpload";

export function HomeView({}): JSX.Element {
  const [file, setFile] = useState<File>();
  const { progress, error, success, upload } = useUpload();

  const handleChange = (event: Event) => {
    const target = event.target as HTMLInputElement;
    if (target.files?.length) {
      setFile(target.files[0]);
    }
  };

  const handleSubmit = () => {
    if (file) {
      upload(file);
    }
  };

  return (
    <div class="space-y-2">
      <label class="cursor-pointer">
        <input
          class="block h-full w-full"
          onChange={handleChange}
          type="file"
          name="File"
        />
      </label>

      <button
        class="rounded bg-slate-300 px-3 py-1.5"
        onClick={handleSubmit}
        type="button"
      >
        Submit
      </button>

      {success ? (
        <p>File uploaded!</p>
      ) : progress ? (
        <p>
          Progress: {progress.current} / {progress.target} (
          {progress.completion * 100}%)
        </p>
      ) : null}

      {error ? <p className="font-bold text-red-400">{error.message}</p> : null}
    </div>
  );
}
