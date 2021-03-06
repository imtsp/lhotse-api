// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: agent/metadata.proto

package com.idefav.lhotse.api.agent;

public final class Metadata {
  private Metadata() {}
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface AgentMetaOrBuilder extends
      // @@protoc_insertion_point(interface_extends:lhotse.api.agent.AgentMeta)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string version = 1 [json_name = "version"];</code>
     * @return The version.
     */
    java.lang.String getVersion();
    /**
     * <code>string version = 1 [json_name = "version"];</code>
     * @return The bytes for version.
     */
    com.google.protobuf.ByteString
        getVersionBytes();
  }
  /**
   * Protobuf type {@code lhotse.api.agent.AgentMeta}
   */
  public static final class AgentMeta extends
      com.google.protobuf.GeneratedMessageV3 implements
      // @@protoc_insertion_point(message_implements:lhotse.api.agent.AgentMeta)
      AgentMetaOrBuilder {
  private static final long serialVersionUID = 0L;
    // Use AgentMeta.newBuilder() to construct.
    private AgentMeta(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
      super(builder);
    }
    private AgentMeta() {
      version_ = "";
    }

    @java.lang.Override
    @SuppressWarnings({"unused"})
    protected java.lang.Object newInstance(
        UnusedPrivateParameter unused) {
      return new AgentMeta();
    }

    @java.lang.Override
    public final com.google.protobuf.UnknownFieldSet
    getUnknownFields() {
      return this.unknownFields;
    }
    private AgentMeta(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      this();
      if (extensionRegistry == null) {
        throw new java.lang.NullPointerException();
      }
      com.google.protobuf.UnknownFieldSet.Builder unknownFields =
          com.google.protobuf.UnknownFieldSet.newBuilder();
      try {
        boolean done = false;
        while (!done) {
          int tag = input.readTag();
          switch (tag) {
            case 0:
              done = true;
              break;
            case 10: {
              java.lang.String s = input.readStringRequireUtf8();

              version_ = s;
              break;
            }
            default: {
              if (!parseUnknownField(
                  input, unknownFields, extensionRegistry, tag)) {
                done = true;
              }
              break;
            }
          }
        }
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        throw e.setUnfinishedMessage(this);
      } catch (java.io.IOException e) {
        throw new com.google.protobuf.InvalidProtocolBufferException(
            e).setUnfinishedMessage(this);
      } finally {
        this.unknownFields = unknownFields.build();
        makeExtensionsImmutable();
      }
    }
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.idefav.lhotse.api.agent.Metadata.internal_static_lhotse_api_agent_AgentMeta_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.idefav.lhotse.api.agent.Metadata.internal_static_lhotse_api_agent_AgentMeta_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.idefav.lhotse.api.agent.Metadata.AgentMeta.class, com.idefav.lhotse.api.agent.Metadata.AgentMeta.Builder.class);
    }

    public static final int VERSION_FIELD_NUMBER = 1;
    private volatile java.lang.Object version_;
    /**
     * <code>string version = 1 [json_name = "version"];</code>
     * @return The version.
     */
    @java.lang.Override
    public java.lang.String getVersion() {
      java.lang.Object ref = version_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs = 
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        version_ = s;
        return s;
      }
    }
    /**
     * <code>string version = 1 [json_name = "version"];</code>
     * @return The bytes for version.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getVersionBytes() {
      java.lang.Object ref = version_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b = 
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        version_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!getVersionBytes().isEmpty()) {
        com.google.protobuf.GeneratedMessageV3.writeString(output, 1, version_);
      }
      unknownFields.writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!getVersionBytes().isEmpty()) {
        size += com.google.protobuf.GeneratedMessageV3.computeStringSize(1, version_);
      }
      size += unknownFields.getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof com.idefav.lhotse.api.agent.Metadata.AgentMeta)) {
        return super.equals(obj);
      }
      com.idefav.lhotse.api.agent.Metadata.AgentMeta other = (com.idefav.lhotse.api.agent.Metadata.AgentMeta) obj;

      if (!getVersion()
          .equals(other.getVersion())) return false;
      if (!unknownFields.equals(other.unknownFields)) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + VERSION_FIELD_NUMBER;
      hash = (53 * hash) + getVersion().hashCode();
      hash = (29 * hash) + unknownFields.hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input);
    }
    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessageV3
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(com.idefav.lhotse.api.agent.Metadata.AgentMeta prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code lhotse.api.agent.AgentMeta}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:lhotse.api.agent.AgentMeta)
        com.idefav.lhotse.api.agent.Metadata.AgentMetaOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return com.idefav.lhotse.api.agent.Metadata.internal_static_lhotse_api_agent_AgentMeta_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return com.idefav.lhotse.api.agent.Metadata.internal_static_lhotse_api_agent_AgentMeta_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                com.idefav.lhotse.api.agent.Metadata.AgentMeta.class, com.idefav.lhotse.api.agent.Metadata.AgentMeta.Builder.class);
      }

      // Construct using com.idefav.lhotse.api.agent.Metadata.AgentMeta.newBuilder()
      private Builder() {
        maybeForceBuilderInitialization();
      }

      private Builder(
          com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
        super(parent);
        maybeForceBuilderInitialization();
      }
      private void maybeForceBuilderInitialization() {
        if (com.google.protobuf.GeneratedMessageV3
                .alwaysUseFieldBuilders) {
        }
      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        version_ = "";

        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return com.idefav.lhotse.api.agent.Metadata.internal_static_lhotse_api_agent_AgentMeta_descriptor;
      }

      @java.lang.Override
      public com.idefav.lhotse.api.agent.Metadata.AgentMeta getDefaultInstanceForType() {
        return com.idefav.lhotse.api.agent.Metadata.AgentMeta.getDefaultInstance();
      }

      @java.lang.Override
      public com.idefav.lhotse.api.agent.Metadata.AgentMeta build() {
        com.idefav.lhotse.api.agent.Metadata.AgentMeta result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public com.idefav.lhotse.api.agent.Metadata.AgentMeta buildPartial() {
        com.idefav.lhotse.api.agent.Metadata.AgentMeta result = new com.idefav.lhotse.api.agent.Metadata.AgentMeta(this);
        result.version_ = version_;
        onBuilt();
        return result;
      }

      @java.lang.Override
      public Builder clone() {
        return super.clone();
      }
      @java.lang.Override
      public Builder setField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.setField(field, value);
      }
      @java.lang.Override
      public Builder clearField(
          com.google.protobuf.Descriptors.FieldDescriptor field) {
        return super.clearField(field);
      }
      @java.lang.Override
      public Builder clearOneof(
          com.google.protobuf.Descriptors.OneofDescriptor oneof) {
        return super.clearOneof(oneof);
      }
      @java.lang.Override
      public Builder setRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          int index, java.lang.Object value) {
        return super.setRepeatedField(field, index, value);
      }
      @java.lang.Override
      public Builder addRepeatedField(
          com.google.protobuf.Descriptors.FieldDescriptor field,
          java.lang.Object value) {
        return super.addRepeatedField(field, value);
      }
      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof com.idefav.lhotse.api.agent.Metadata.AgentMeta) {
          return mergeFrom((com.idefav.lhotse.api.agent.Metadata.AgentMeta)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(com.idefav.lhotse.api.agent.Metadata.AgentMeta other) {
        if (other == com.idefav.lhotse.api.agent.Metadata.AgentMeta.getDefaultInstance()) return this;
        if (!other.getVersion().isEmpty()) {
          version_ = other.version_;
          onChanged();
        }
        this.mergeUnknownFields(other.unknownFields);
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        com.idefav.lhotse.api.agent.Metadata.AgentMeta parsedMessage = null;
        try {
          parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          parsedMessage = (com.idefav.lhotse.api.agent.Metadata.AgentMeta) e.getUnfinishedMessage();
          throw e.unwrapIOException();
        } finally {
          if (parsedMessage != null) {
            mergeFrom(parsedMessage);
          }
        }
        return this;
      }

      private java.lang.Object version_ = "";
      /**
       * <code>string version = 1 [json_name = "version"];</code>
       * @return The version.
       */
      public java.lang.String getVersion() {
        java.lang.Object ref = version_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          version_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string version = 1 [json_name = "version"];</code>
       * @return The bytes for version.
       */
      public com.google.protobuf.ByteString
          getVersionBytes() {
        java.lang.Object ref = version_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b = 
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          version_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string version = 1 [json_name = "version"];</code>
       * @param value The version to set.
       * @return This builder for chaining.
       */
      public Builder setVersion(
          java.lang.String value) {
        if (value == null) {
    throw new NullPointerException();
  }
  
        version_ = value;
        onChanged();
        return this;
      }
      /**
       * <code>string version = 1 [json_name = "version"];</code>
       * @return This builder for chaining.
       */
      public Builder clearVersion() {
        
        version_ = getDefaultInstance().getVersion();
        onChanged();
        return this;
      }
      /**
       * <code>string version = 1 [json_name = "version"];</code>
       * @param value The bytes for version to set.
       * @return This builder for chaining.
       */
      public Builder setVersionBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
        
        version_ = value;
        onChanged();
        return this;
      }
      @java.lang.Override
      public final Builder setUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.setUnknownFields(unknownFields);
      }

      @java.lang.Override
      public final Builder mergeUnknownFields(
          final com.google.protobuf.UnknownFieldSet unknownFields) {
        return super.mergeUnknownFields(unknownFields);
      }


      // @@protoc_insertion_point(builder_scope:lhotse.api.agent.AgentMeta)
    }

    // @@protoc_insertion_point(class_scope:lhotse.api.agent.AgentMeta)
    private static final com.idefav.lhotse.api.agent.Metadata.AgentMeta DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new com.idefav.lhotse.api.agent.Metadata.AgentMeta();
    }

    public static com.idefav.lhotse.api.agent.Metadata.AgentMeta getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<AgentMeta>
        PARSER = new com.google.protobuf.AbstractParser<AgentMeta>() {
      @java.lang.Override
      public AgentMeta parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        return new AgentMeta(input, extensionRegistry);
      }
    };

    public static com.google.protobuf.Parser<AgentMeta> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<AgentMeta> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public com.idefav.lhotse.api.agent.Metadata.AgentMeta getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_lhotse_api_agent_AgentMeta_descriptor;
  private static final 
    com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internal_static_lhotse_api_agent_AgentMeta_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\024agent/metadata.proto\022\020lhotse.api.agent" +
      "\"%\n\tAgentMeta\022\030\n\007version\030\001 \001(\tR\007versionB" +
      "/\n\033com.idefav.lhotse.api.agentZ\020lhotse/a" +
      "pi/agentb\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_lhotse_api_agent_AgentMeta_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_lhotse_api_agent_AgentMeta_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessageV3.FieldAccessorTable(
        internal_static_lhotse_api_agent_AgentMeta_descriptor,
        new java.lang.String[] { "Version", });
  }

  // @@protoc_insertion_point(outer_class_scope)
}
