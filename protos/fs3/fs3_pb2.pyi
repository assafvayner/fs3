from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
GREAT_SUCCESS: Status
ILLEGAL_PATH: Status
INTERNAL_ERROR: Status
NOT_FOUND: Status

class CopyReply(_message.Message):
    __slots__ = ["file_path", "status"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    status: Status
    def __init__(self, status: _Optional[_Union[Status, str]] = ..., file_path: _Optional[str] = ...) -> None: ...

class CopyRequest(_message.Message):
    __slots__ = ["file_content", "file_path"]
    FILE_CONTENT_FIELD_NUMBER: _ClassVar[int]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_content: bytes
    file_path: str
    def __init__(self, file_path: _Optional[str] = ..., file_content: _Optional[bytes] = ...) -> None: ...

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
    __slots__ = ["file_path"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    def __init__(self, file_path: _Optional[str] = ...) -> None: ...

class RemoveReply(_message.Message):
    __slots__ = ["file_path", "status"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    status: Status
    def __init__(self, status: _Optional[_Union[Status, str]] = ..., file_path: _Optional[str] = ...) -> None: ...

class RemoveRequest(_message.Message):
    __slots__ = ["file_path"]
    FILE_PATH_FIELD_NUMBER: _ClassVar[int]
    file_path: str
    def __init__(self, file_path: _Optional[str] = ...) -> None: ...

class Status(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
