import requests
import traceback
import json
from typing import List, Dict, Union
from functools import lru_cache
from dataclasses import dataclass


@dataclass(frozen=True)
class PistonResponse:
    """
    Piston API response class.
    """

    language: str
    exit_code: int
    output: str


class CodeExecutor:
    """
    Base class for the remote code execution.
    """

    def __init__(self) -> None:
        self.runtimes = self.__get_runtimes()

    @staticmethod
    def __get_runtimes() -> List[Dict[str, Union[str, List[str]]]]:
        """
        Returns a list of all available runtimes offered by the piston api.
        """
        try:
            runtime_url = "https://emkc.org/api/v2/piston/runtimes"
            r = requests.get(runtime_url)
            data = json.loads(r.text)
            runtimes: List[Dict[str, Union[str, List[str]]]] = []
            for langs in data:
                runtimes.append(
                    {
                        "language": langs["language"],
                        "version": langs["version"],
                        "aliases": [i for i in langs["aliases"]],
                    }
                )

            return runtimes

        except Exception as e:
            traceback.print_exception(e)
            return []

    @lru_cache(maxsize=50)
    def execute_code(self, language: str, code: str) -> PistonResponse:
        """
        Executes the given code in the given language and version.
        """
        try:
            execute_url = "https://emkc.org/api/v2/piston/execute"
            all_langs = [runtime["language"] for runtime in self.runtimes]
            _aliases = [runtime["aliases"] for runtime in self.runtimes]
            aliases = [i for sublist in _aliases for i in sublist]

            if language not in all_langs and language not in aliases:
                return PistonResponse(
                    language, -1, f"Language {language} is not supported."
                )

            for runtime in self.runtimes:
                if runtime["language"] == language or language in runtime["aliases"]:
                    version = runtime["version"]
                    break

            else:
                return PistonResponse(
                    language, -1, f"Language {language} is not supported."
                )

            payload = {
                "language": language,
                "version": version,
                "files": [
                    {
                        "name": "prog",
                        "content": code,
                    }
                ],
            }

            resp = requests.post(execute_url, json=payload)

            if resp.status_code == 200:
                data = json.loads(resp.text)

                run_data = data["run"]

                return PistonResponse(language, run_data["code"], run_data["output"])

            else:
                return PistonResponse(
                    language, -1, f"Internal error: {resp.status_code}"
                )

        except Exception as e:
            traceback.print_exception(e)
            return PistonResponse(language, -1, "Couldn't connect at the moment.")


if __name__ == "__main__":
    rce = CodeExecutor()
    # print(rce.runtimes)
    print(rce.execute_code(language="vlang", code="print('hello')"))
    print(rce.execute_code(language="vlang", code="print('hello')"))
