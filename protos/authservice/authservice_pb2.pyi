from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetNewTokenReply(_message.Message):
    __slots__ = ["status", "token", "username"]
    class Status(_message.Message):
        __slots__ = ["message", "success"]
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        SUCCESS_FIELD_NUMBER: _ClassVar[int]
        message: str
        success: bool
        def __init__(self, success: bool = ..., message: _Optional[str] = ...) -> None: ...
    STATUS_FIELD_NUMBER: _ClassVar[int]
    TOKEN_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    status: GetNewTokenReply.Status
    token: str
    username: str
    def __init__(self, username: _Optional[str] = ..., status: _Optional[_Union[GetNewTokenReply.Status, _Mapping]] = ..., token: _Optional[str] = ...) -> None: ...

class GetNewTokenRequest(_message.Message):
    __slots__ = ["password", "previous_token", "username"]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    PREVIOUS_TOKEN_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    password: str
    previous_token: str
    username: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ..., previous_token: _Optional[str] = ...) -> None: ...

class NewUserReply(_message.Message):
    __slots__ = ["status", "username"]
    class Status(_message.Message):
        __slots__ = ["message", "success"]
        MESSAGE_FIELD_NUMBER: _ClassVar[int]
        SUCCESS_FIELD_NUMBER: _ClassVar[int]
        message: str
        success: bool
        def __init__(self, success: bool = ..., message: _Optional[str] = ...) -> None: ...
    STATUS_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    status: NewUserReply.Status
    username: str
    def __init__(self, username: _Optional[str] = ..., status: _Optional[_Union[NewUserReply.Status, _Mapping]] = ...) -> None: ...

class NewUserRequest(_message.Message):
    __slots__ = ["password", "username"]
    PASSWORD_FIELD_NUMBER: _ClassVar[int]
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    password: str
    username: str
    def __init__(self, username: _Optional[str] = ..., password: _Optional[str] = ...) -> None: ...
