from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

BAD_REQUEST: Status
DESCRIPTOR: _descriptor.FileDescriptor
GREAT_SUCCESS: Status
ILLEGAL_PATH: Status
INTERNAL_ERROR: Status
NOT_FOUND: Status
UNSPECIFIED: Status

class CopyReply(_message.Message):
    __slots__ = ["file_path", "status"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    status: Status
    def __init__(self, status: _Optional[_Union[Status, str]] = ..., file_path: _Optional[str] = ...) -> None: ...

class CopyRequest(_message.Message):
    __slots__ = ["file_content", "file_path", "token"]
    FILE_CONTENT_FIELD_NUMBER: _ClassVar[int]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    file_content: bytes
    file_path: str
    token: str
    def __init__(self, file_path: _Optional[str] = ..., file_content: _Optional[bytes] = ..., token: _Optional[str] = ...) -> None: ...

class DescribeReply(_message.Message):
    __slots__ = ["directory", "file", "path", "status"]
    class Directory(_message.Message):
        __slots__ = ["directoryname", "files", "subdirectories"]
        DIRECTORYNAME_FIELD_NUMBER: _ClassVar[int]
        FILES_FIELD_NUMBER: _ClassVar[int]
        SUBDIRECTORIES_FIELD_NUMBER: _ClassVar[int]
        directoryname: str
        files: _containers.RepeatedScalarFieldContainer[str]
        subdirectories: _containers.RepeatedScalarFieldContainer[str]
        def __init__(self, directoryname: _Optional[str] = ..., files: _Optional[_Iterable[str]] = ..., subdirectories: _Optional[_Iterable[str]] = ...) -> None: ...
    class File(_message.Message):
        __slots__ = ["filename"]
        FILENAME_FIELD_NUMBER: _ClassVar[int]
        filename: str
        def __init__(self, filename: _Optional[str] = ...) -> None: ...
    DIRECTORY_FIELD_NUMBER: _ClassVar[int]
    FILE_FIELD_NUMBER: _ClassVar[int]
    PATH_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    directory: DescribeReply.Directory
    file: DescribeReply.File
    path: str
    status: Status
    def __init__(self, status: _Optional[_Union[Status, str]] = ..., path: _Optional[str] = ..., file: _Optional[_Union[DescribeReply.File, _Mapping]] = ..., directory: _Optional[_Union[DescribeReply.Directory, _Mapping]] = ...) -> None: ...

class DescribeRequest(_message.Message):
    __slots__ = ["path", "token"]
    PATH_FIELD_NUMBER: _ClassVar[int]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    path: str
    token: str
    def __init__(self, path: _Optional[str] = ..., token: _Optional[str] = ...) -> None: ...

class GetReply(_message.Message):
    __slots__ = ["file_content", "file_path", "status"]
    FILE_CONTENT_FIELD_NUMBER: _ClassVar[int]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    file_content: bytes
    file_path: str
    status: Status
    def __init__(self, status: _Optional[_Union[Status, str]] = ..., file_path: _Optional[str] = ..., file_content: _Optional[bytes] = ...) -> None: ...

class GetRequest(_message.Message):
    __slots__ = ["file_path", "token"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    token: str
    def __init__(self, file_path: _Optional[str] = ..., token: _Optional[str] = ...) -> None: ...

class RemoveReply(_message.Message):
    __slots__ = ["file_path", "status"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    status: Status
    def __init__(self, status: _Optional[_Union[Status, str]] = ..., file_path: _Optional[str] = ...) -> None: ...

class RemoveRequest(_message.Message):
    __slots__ = ["file_path", "token"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    token: str
    def __init__(self, file_path: _Optional[str] = ..., token: _Optional[str] = ...) -> None: ...

class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
